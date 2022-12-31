pipeline{
    agent any
    tools {
        go 'Go 1.19.3'
    }
    environment {
        // GO114MODULE = 'on'
        // CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages{
        stage("Unit Testing"){
            steps{
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'make unit-tests'
            }
            post{
                success{
                    echo "========UNIT TEST EXECUTION COMPLETED========"
                }
                failure{
                    echo "========UNIT TEST EXECUTION FAILED========"
                }
            }
        }
        stage("Build"){
            steps{
                
                echo 'BUILD EXECUTION STARTED'
                sh 'make build'
            }
            post{
                success{
                    echo "========BUILD EXECUTION COMPLETED========"
                }
                failure{
                    echo "========BUILD EXECUTION FAILED========"
                }
            }
        }
        stage("Deploy"){
            steps{
                echo 'DEPLOYING EXECUTION STARTED'
                sh 'make run-build'
            }
            post{
                success{
                    echo "========DEPLOYING EXECUTION COMPLETED========"
                }
                failure{
                    echo "========DEPLOYING EXECUTION FAILED========"
                }
            }
        }
    }
    post{
        success{
            echo "========pipeline executed successfully ========"
        }
        failure{
            echo "========pipeline execution failed========"
        }
          always {
            emailext body: "${currentBuild.currentResult}: Job ${env.JOB_NAME} build ${env.BUILD_NUMBER}\n More info at: ${env.BUILD_URL}",
                recipientProviders: [[$class: 'DevelopersRecipientProvider'], [$class: 'RequesterRecipientProvider']],
                to: "${params.RECIPIENTS}",
                subject: "Jenkins Build ${currentBuild.currentResult}: Job ${env.JOB_NAME}"

            slackSend channel: '#jenkins', message: "Hi, Do well to find pipeiine status for \
            ${currentBuild.currentResult}\
            Job Name: ${env.JOB_NAME}\
            Build Number: ${env.BUILD_NUMBER}\
            Url: ${env.BUILD_URL}"
            
        }
    }
}
