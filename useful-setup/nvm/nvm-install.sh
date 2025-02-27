#source: https://nodejs.org/en/download/package-manager/current

# installs nvm (Node Version Manager)
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.0/install.sh | bash

# download and install Node.js (you may need to restart the terminal)
nvm install 22

# verifies the right Node.js version is in the environment
node -v # should print `v22.9.0`

# verifies the right npm version is in the environment
npm -v # should print `10.8.3`