node {
   def err = null
   def project = "helm-annotate"
  
   stage("Checkout code") {
     checkout scm
   }
   stage "Run CI script"
     try {
       sh "scripts/ci.sh ${project}"
     } catch (Exception e) {
       currentBuild.result = "FAILURE"
       err = e
     }

   if (err) {
     throw err
   }
}
