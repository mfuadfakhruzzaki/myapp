pipeline {
    agent any

    environment {
        // Pastikan Golang ada di PATH
        PATH = "/usr/local/go/bin:${env.PATH}"
        DEPLOY_DIR = "/var/www/myapp"
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build Backend') {
            steps {
                dir('backend') {
                    echo 'Building Backend...'
                    sh 'go mod tidy'
                    sh 'go build -o myapp-backend'
                }
            }
        }
        stage('Test Backend') {
            steps {
                dir('backend') {
                    echo 'Running Backend Tests...'
                    sh 'go test ./...'
                }
            }
        }
        stage('Build Frontend') {
            steps {
                dir('frontend') {
                    echo 'Installing frontend dependencies...'
                    sh 'npm install'
                    echo 'Building Frontend...'
                    sh 'npm run build'
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying application...'
                // Deploy Backend:
                dir('backend') {
                    echo 'Deploying Backend...'
                    sh 'nohup ./myapp-backend > backend.log 2>&1 &'
                }
                // Deploy Frontend:
                dir('frontend') {
                    echo 'Deploying Frontend...'
                    // Pastikan direktori target sudah ada secara manual dengan izin yang tepat
                    sh 'cp -r dist/* ${DEPLOY_DIR}/'
                }
            }
        }
    }
    post {
        success {
            echo 'Pipeline executed successfully!'
        }
        failure {
            echo 'Pipeline execution failed!'
        }
    }
}
