pipeline {
    agent any

    environment {
        // Pastikan direktori Golang sudah ditambahkan ke PATH
        PATH = "/usr/local/go/bin:${env.PATH}"
    }

    stages {
        stage('Checkout') {
            steps {
                // Mengambil kode dari repository yang telah dikonfigurasi di job Jenkins
                checkout scm
            }
        }
        stage('Build Backend') {
            steps {
                dir('backend') {
                    // Pastikan dependency terpasang dan build binary backend
                    sh 'go mod tidy'
                    sh 'go build -o myapp-backend'
                }
            }
        }
        stage('Test Backend') {
            steps {
                dir('backend') {
                    // Jalankan unit test pada backend
                    sh 'go test ./...'
                }
            }
        }
        stage('Build Frontend') {
            steps {
                dir('frontend') {
                    // Pasang dependency dan build aplikasi frontend menggunakan Vite
                    sh 'npm install'
                    sh 'npm run build'
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying application...'
                // Deploy Backend:
                // Opsi: Matikan instance backend yang lama jika diperlukan (tidak ditunjukkan di sini)
                dir('backend') {
                    // Menjalankan binary backend di background. Hasil output akan dialihkan ke file log.
                    sh 'nohup ./myapp-backend > backend.log 2>&1 &'
                }
                // Deploy Frontend:
                dir('frontend') {
                    // Pastikan direktori target (/var/www/myapp/) sudah ada dan dapat diakses oleh Jenkins.
                    // Menyalin hasil build (misalnya, folder "dist") ke direktori deploy.
                    sh 'cp -r dist/* /var/www/myapp/'
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline executed successfully!'
            // Anda bisa menambahkan notifikasi di sini (misalnya, email atau Slack)
        }
        failure {
            echo 'Pipeline execution failed!'
            // Tambahkan langkah notifikasi jika diperlukan
        }
    }
}
