pipeline{
    agent{
        label "API Latencies"
    }
    tools {
        go 'go1.19.3'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
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
    }
    stages{
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
    }
    stages{
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
    }
}