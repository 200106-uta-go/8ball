pipeline {
   agent any

   tools {
      go 'go1.13'
   }

   stages {
      stage('Pull') {
          steps {
              git 'https://github.com/200106-uta-go/8ball.git'   
          }
      }
      stage('Test') {
          steps {
              sh "go test ./..."
          }
      }
      stage('Build') {
         steps {
            sh "go build cmd/8ball/*.go"
            sh "go build cmd/8balld/*.go"
         }

         post {
            success {
               sh "echo SUCCESS"
            }
         }
      }
   }
}
