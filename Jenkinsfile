pipeline {
    agent any
    tools { go 'go1.19' }
    stages {
        stage('versionCheck') {
            steps {
                sh 'go version'
            }
        }
        stage('envCheck') {
            steps {
                sh 'env | sort'
            }
        }
        stage('build ') {
            steps {
                sh 'GOOS=linux GOARCH=amd64 go build -o ${WORKSPACE}/bin/linux/amd64/eagerLaser .'
                sh 'GOOS=linux GOARCH=arm64 go build -o ${WORKSPACE}/bin/linux/arm64/eagerLaser .'
                sh 'GOOS=darwin GOARCH=arm64 go build -o ${WORKSPACE}/bin/darwin/arm64/eagerLaser .'
                sh 'GOOS=darwin GOARCH=amd64 go build -o ${WORKSPACE}/bin/darwin/amd64/eagerLaser .'
            }
        }
    }
    post {
        success {
            archiveArtifacts artifacts: '${WORKSPACE}/bin/**/*', fingerprint: true
        }
    }
}