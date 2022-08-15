pipeline {
    agent { any }
    tools { go 'go1.19' }
    stages {
        stage('build') {
            steps {
                sh 'go version'
            }
        }
    }
}