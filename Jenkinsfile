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
            script {
                def status = currentBuild.currentResult
                def state = status == 'SUCCESS' ? 'success' : 'failure'
                def description = "Build ${status.toLowerCase()}"

                sh """
                    curl -X POST \\
                         -H "Authorization: token \${GITHUB_TOKEN}" \\
                         -H "Content-Type: application/json" \\
                         -d '{
                             "state": "${state}",
                             "target_url": "${env.BUILD_URL}",
                             "description": "${description}",
                             "context": "continuous-integration/jenkins"
                         }' \\
                         "https://api.github.com/repos/\${GITHUB_REPO}/statuses/\${GIT_COMMIT}"
                """
            }
        }
    }
}
