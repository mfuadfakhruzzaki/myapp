pipeline {
    agent any

    environment {
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
                    // Debug: Tampilkan file yang dihasilkan
                    sh 'ls -l'
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
                    // Jalankan binary backend secara background dengan input dialihkan
                    sh 'nohup ./myapp-backend > backend.log 2>&1 < /dev/null &'
                    // Debug: Tampilkan proses yang berjalan
                    sh 'ps aux | grep myapp-backend'
                }
                // Deploy Frontend:
                dir('frontend') {
                    echo 'Deploying Frontend...'
                    sh 'mkdir -p ${DEPLOY_DIR}'
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
