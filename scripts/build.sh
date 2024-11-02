#!/bin/sh

IMAGE_NAME="test"

DOCKERFILE_PATH="../docker/Dockerfile"

# Выполняем команду сборки Docker
docker build -t "$IMAGE_NAME" -f "$DOCKERFILE_PATH" ..dock

# Проверяем успешность сборки
if [ $? -eq 0 ]; then
  echo "Docker image '$IMAGE_NAME' built successfully."
else
  echo "Failed to build Docker image '$IMAGE_NAME'."
  exit 1
fi
