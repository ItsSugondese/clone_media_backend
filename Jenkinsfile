pipeline {
    agent any   // run on any Jenkins agent

    stages {

        stage('Build') {
            steps {
                echo "Building project..."
                sh 'go build .'
            }
        }
    }

    post {
        success {
            genericstatus([
                description: 'Build succeeded',
                state: 'success',
                targetUrl: env.BUILD_URL
            ])
        }
        failure {
            genericstatus([
                description: 'Build failed',
                state: 'failure',
                targetUrl: env.BUILD_URL
            ])
        }
    }
}
