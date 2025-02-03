pipeline {
    agent any

    environment {
        // Pastikan direktori Golang ada di PATH
        PATH = "/usr/local/go/bin:${env.PATH}"
        // Jika diperlukan, Anda bisa menambahkan variabel lain, misalnya target deployment
        DEPLOY_DIR = "/var/www/myapp"
    }

    stages {
        stage('Checkout') {
            steps {
                // Mengambil kode dari repository
                checkout scm
            }
        }
        stage('Build Backend') {
            steps {
                dir('backend') {
                    echo 'Building Backend...'
                    // Pastikan dependency terinstall
                    sh 'go mod tidy'
                    // Build binary backend
                    sh 'go build -o myapp-backend'
                }
            }
        }
        stage('Test Backend') {
            steps {
                dir('backend') {
                    echo 'Running Backend Tests...'
                    // Jalankan unit test. Jika tidak ada test, stage ini akan selesai tanpa output.
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
                    // Jika diperlukan, tambahkan logika untuk menghentikan instance backend yang lama
                    // Jalankan backend sebagai background process menggunakan nohup.
                    sh 'nohup ./myapp-backend > backend.log 2>&1 &'
                }
                // Deploy Frontend:
                dir('frontend') {
                    echo 'Deploying Frontend...'
                    // Buat direktori target jika belum ada
                    sh 'mkdir -p ${DEPLOY_DIR}'
                    // Salin seluruh isi folder "dist" ke direktori target
                    sh 'cp -r dist/* ${DEPLOY_DIR}/'
                }
            }
        }
    }
    post {
        success {
            echo 'Pipeline executed successfully!'
            // Di sini Anda bisa menambahkan notifikasi, misalnya mengirim email atau Slack message.
        }
        failure {
            echo 'Pipeline execution failed!'
            // Tambahkan langkah notifikasi error jika diperlukan.
        }
    }
}
