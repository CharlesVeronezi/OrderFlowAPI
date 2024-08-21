import http from 'k6/http'

export default function () {
    const url = 'http://localhost:9999/users';
    const payload = JSON.stringify({
        us_firstname: "John",
        us_lastname: "Snow",
        us_email: "teste@gmail.com",
        us_vbactive: true
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}