FROM python:3.10-slim

# Update and install dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    cmake \
    && rm -rf /var/lib/apt/lists/* # Clean up apt cache to reduce image size

# Install the Python package
RUN pip3 install llama-cpp-python[server]

# Expose the ports
EXPOSE 8000

# Make a directory for the models to be mounted into
RUN mkdir -p /var/lib/models

# Set the environment 
ENV MODEL "/var/lib/models/ggjt-model.bin"

# Start the server
CMD ["python3", "-m", "llama_cpp.server"]
