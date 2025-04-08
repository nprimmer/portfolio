# Use an official Node.js runtime as the base image
FROM node:19

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application to the working directory
COPY . .

# Build the app
RUN npm run build

# Use a lightweight node image for the built app
FROM node:19-slim

# Set the working directory
WORKDIR /usr/src/app

# Copy the build folder from the previous stage to the new working directory
COPY --from=0 /usr/src/app/build ./build

# Install serve to serve the built React app
RUN npm install -g serve

# Expose the port the app will run on
EXPOSE 5000

# Command to run the application
CMD ["serve", "-s", "build"]
