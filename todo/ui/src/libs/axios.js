import axios from 'axios';

export default axios.create({
    // For dev running from docker-compose
    // baseURL: "http://localhost:8080",
    // For prod talking to API in Kubernetes via Ingress
    baseURL: "/api",
});
