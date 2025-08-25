pipeline {
    agent any   // run on any Jenkins agent

    stages {
        stage('Checkout') {
            steps {
                // Pull code from GitHub
                git branch: 'main', url: 'https://github.com/ItsSugondese/clone_media_backend.git'
            }
        }

        stage('Build') {
            steps {
                echo "Building project..."
                sh 'go build .'
            }
        }

        stage('Test') {
            steps {
                echo "Running tests..."
                sh 'go test ./...'
            }
        }

        stage('Package') {
            steps {
                echo "Packaging build..."
                sh 'tar -czf build.tar.gz *'
            }
        }
    }
}
