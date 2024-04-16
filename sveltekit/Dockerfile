# Use the official Node.js 14 image as the base image
FROM node:21

# Set the working directory inside the container
WORKDIR /app

COPY . .

RUN npm ci
RUN npm install

# Build the SvelteKit project
RUN npm run build

# Expose the port that the SvelteKit app will run on
EXPOSE 3000

# Start the SvelteKit app
CMD ["node", "build"]

