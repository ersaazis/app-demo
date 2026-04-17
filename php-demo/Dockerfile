FROM php:8.2-alpine

# Set working directory
WORKDIR /var/www/html

# Install system dependencies
RUN apk add --no-cache \
    curl \
    libpng-dev \
    libxml2-dev \
    libzip-dev \
    zip \
    unzip \
    oniguruma-dev \
    icu-dev \
    git \
    bash

# Install PHP extensions
RUN docker-php-ext-install \
    bcmath \
    gd \
    mbstring \
    pdo \
    pdo_mysql \
    zip \
    intl

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Allow unlimited memory for Composer
ENV COMPOSER_MEMORY_LIMIT=-1

# Copy composer files first (cache optimization)
COPY composer.json composer.lock ./

# Install PHP dependencies
RUN composer install \
    --no-dev \
    --optimize-autoloader \
    --no-scripts \
    --no-interaction \
    --prefer-dist

# Copy the rest of the application
COPY . .

# Fix permissions (Lumen/Laravel)
RUN chown -R www-data:www-data /var/www/html/storage

# Expose port
EXPOSE 8000

# Run app
CMD ["php", "-S", "0.0.0.0:8000", "-t", "public"]