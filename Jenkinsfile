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
                echo 'Build succeeded - GitHub will receive SUCCESS status'
                // Jenkins automatically sends: {"state": "success", "description": "Build succeeded"}
            }
            failure {
                echo 'Build failed - GitHub will receive FAILURE status'
                // Jenkins automatically sends: {"state": "failure", "description": "Build failed"}
            }
            unstable {
                echo 'Build unstable - GitHub will receive FAILURE status'
                // Jenkins automatically sends: {"state": "failure", "description": "Build unstable"}
            }
            aborted {
                echo 'Build aborted - GitHub will receive ERROR status'
                // Jenkins automatically sends: {"state": "error", "description": "Build aborted"}
            }
        }
}
