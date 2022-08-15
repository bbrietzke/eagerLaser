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
                sh 'TMPDIR=${WORKSPACE_TMP} GOOS=linux GOARCH=amd64 go build -o bin/linux/amd64/eagerLaser .'
                sh 'TMPDIR=${WORKSPACE_TMP} GOOS=linux GOARCH=arm64 go build -o bin/linux/arm64/eagerLaser .'
                sh 'TMPDIR=${WORKSPACE_TMP} GOOS=darwin GOARCH=arm64 go build -o bin/darwin/arm64/eagerLaser .'
                sh 'TMPDIR=${WORKSPACE_TMP} GOOS=darwin GOARCH=amd64 go build -o bin/darwin/amd64/eagerLaser .'
            }
        }
    }
    post {
        success {
            archiveArtifacts artifacts: 'bin/**/*', fingerprint: true
        }
    }
}