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
}
