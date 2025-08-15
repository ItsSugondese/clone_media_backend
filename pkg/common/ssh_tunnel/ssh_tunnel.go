package ssh_tunnel

import (
	"clone_media/constants/env"
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func ConnectWithSSHTunnel() {
	mode := os.Getenv("DB_MODE") // "ssh" or "local"
	if mode == "ssh" {
		err := startSSHTunnel(
			"ec2-user",
			os.Getenv(env.PPM_PATH),
			os.Getenv(env.SERVER_URL),
			"5433",           // local port
			"localhost:5432", // remote Postgres port
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func startSSHTunnel(user, privateKeyPath, sshHost, localPort, remoteHost string) error {
	key, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return err
	}

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // for testing only
	}

	sshConn, err := ssh.Dial("tcp", sshHost+":22", sshConfig)
	if err != nil {
		return err
	}

	listener, err := net.Listen("tcp", "127.0.0.1:"+localPort)
	if err != nil {
		return err
	}

	go func() {
		for {
			localConn, err := listener.Accept()
			if err != nil {
				log.Println("Listener accept error:", err)
				continue
			}

			remoteConn, err := sshConn.Dial("tcp", remoteHost)
			if err != nil {
				log.Println("SSH dial error:", err)
				localConn.Close()
				continue
			}

			go func() {
				defer localConn.Close()
				defer remoteConn.Close()
				go io.Copy(remoteConn, localConn)
				io.Copy(localConn, remoteConn)
			}()
		}
	}()

	fmt.Println("SSH tunnel established on 127.0.0.1:" + localPort)
	return nil
}
