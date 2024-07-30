import http from 'k6/http'

export default function () {
    const url = 'http://localhost:3000/address';
    const payload = JSON.stringify({
        ad_street: "123 Main St",
        ad_city: "Anytown",
        ad_state: "CA",
        ad_zip: "12345",
        ad_country: "USA"
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}