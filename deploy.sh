#!/bin/bash

# Set your Google Cloud project ID
PROJECT_ID="your-project-id"
SERVICE_NAME="davidtang-ai"
REGION="us-central1"

# Build the Docker image
echo "Building Docker image..."
docker build -t gcr.io/${PROJECT_ID}/${SERVICE_NAME} .

# Push to Google Container Registry
echo "Pushing to Google Container Registry..."
docker push gcr.io/${PROJECT_ID}/${SERVICE_NAME}

# Deploy to Cloud Run
echo "Deploying to Cloud Run..."
gcloud run deploy ${SERVICE_NAME} \
  --image gcr.io/${PROJECT_ID}/${SERVICE_NAME} \
  --platform managed \
  --region ${REGION} \
  --allow-unauthenticated \
  --port 8080 \
  --memory 128Mi \
  --cpu 0.083 \
  --max-instances 10 \
  --execution-environment gen1

echo "Deployment complete!"