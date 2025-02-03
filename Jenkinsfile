pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build Backend') {
            steps {
                dir('backend') {
                    sh 'go mod tidy'
                    sh 'go build -o myapp-backend'
                }
            }
        }
        stage('Test Backend') {
            steps {
                dir('backend') {
                    sh 'go test ./...'
                }
            }
        }
        stage('Build Frontend') {
            steps {
                dir('frontend') {
                    sh 'npm install'
                    sh 'npm run build'
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deployment stage: Sesuaikan langkah deployment sesuai kebutuhan Anda.'
            }
        }
    }
    post {
        success {
            echo 'Pipeline berhasil dijalankan!'
        }
        failure {
            echo 'Pipeline gagal. Periksa log untuk detail error.'
        }
    }
}
