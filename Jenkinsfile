
pipeline {
    agent any

    options {
        timestamps()
    }

    stages {
        stage('Checkout') {
            steps {
                echo 'Source code checked out from GitHub'
            }
        }

        stage('Verify Go') {
            steps {
                sh 'go version'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o go-app .'
            }
        }

        stage('Unit Tests') {
            steps {
                sh 'go test -v ./...'
            }
        }

        stage('Static Analysis') {
            steps {
                sh 'go vet ./...'
            }
        }
    }

    post {
        success {
            echo 'CI Pipeline completed successfully!'
        }

        failure {
            echo 'CI Pipeline failed. Check console output.'
        }
    }
}
