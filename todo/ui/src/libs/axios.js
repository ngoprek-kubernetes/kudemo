import axios from 'axios';

export default axios.create({
    // For dev running from docker-compose
    baseURL: "http://localhost:8080",
});
