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
            always {
                // This will report the build status back to GitHub
                publishHTML([
                    allowMissing: false,
                    alwaysLinkToLastBuild: true,
                    keepAll: true,
                    reportDir: 'reports',
                    reportFiles: 'index.html',
                    reportName: 'Build Report'
                ])
            }
        }
}
