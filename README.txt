Steps for setup

Obtain both linux binaries for CPSCTaipei and the logger program.
Create private key for the testing server, and deploy the public key on all other servers. The testing server is the server that runs both the tracker, logger and seeder. It is where you will run the deployment bash script.

Run:
mkdir ~/cs416/test 
cp CPSC416Taipei logger deployment.bash ~/cs416/test
cp yourfile ~/cs416/test

Once this is done, run the deployment.bash with the working directory being ~/cs416/test on the wanted machine. For our tests we ssh'ed into remote.ugrad.cs.ubc.ca and ran ./deployment.bash with a 100MB file. The script should prompt you for the path to file you want to distribute. After running, you should have 4 trials of data for each algorithm. 
