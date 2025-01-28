# Utiliser une image Node.js pour construire le frontend
FROM node:18 as builder

# Définir le dossier de travail dans le conteneur
WORKDIR /app

# Copier les fichiers du frontend
COPY package.json package-lock.json ./
RUN npm install

# Copier le reste du code source
COPY . .

# Construire le projet React
RUN npm run build

# Image légère pour servir les fichiers construits
FROM nginx:alpine

# Copier les fichiers construits dans le dossier nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# Exposer le port utilisé par nginx
EXPOSE 80

# Commande par défaut pour démarrer nginx
CMD ["nginx", "-g", "daemon off;"]
