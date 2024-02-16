# Use the official Node.js 14 image as the base image
FROM node:14

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the project files to the working directory
COPY . .

# Build the SvelteKit project
# RUN npm run build

# Expose the port that the SvelteKit app will run on
EXPOSE 5000

# Start the SvelteKit app
CMD ["npm", "run", "start"]
