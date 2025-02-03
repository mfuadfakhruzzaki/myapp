pipeline {
    agent any

    environment {
        PATH = "/usr/local/go/bin:${env.PATH}"
        DEPLOY_DIR = "/var/www/myapp"
        // Environment variables untuk koneksi database
        DB_HOST = "127.0.0.1"
        DB_PORT = "5432"
        DB_USER = "postgres"
        DB_PASSWORD = "020803"
        DB_NAME = "myappdb"
    }

    triggers {
        // Poll SCM setiap 5 menit untuk memicu pipeline saat ada perubahan
        pollSCM('H/5 * * * *')
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
        stage('Restart Service') {
            steps {
                echo 'Reloading daemon and restarting backend service...'
                sh 'sudo systemctl daemon-reload'
                sh 'sudo systemctl restart myapp-backend.service'
                // Debug: Tampilkan status service backend
                sh 'sudo systemctl status myapp-backend.service'
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
