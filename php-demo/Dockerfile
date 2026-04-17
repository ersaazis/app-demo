FROM php:8.1-alpine

# Set working directory
WORKDIR /var/www/html

# Install system dependencies if needed
RUN apk add --no-cache \
    curl \
    libpng-dev \
    libxml2-dev \
    zip \
    unzip

# Install PHP extensions
# Note: Lumen needs some extensions, usually included in alpine or easy to add
RUN docker-php-ext-install bcmath gd

# Copy application files
COPY . .

# Expose port
EXPOSE 8000

# Start PHP built-in server
CMD ["php", "-S", "0.0.0.0:8000", "-t", "public"]
