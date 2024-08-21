import http from 'k6/http'

export default function () {
    const url = 'http://192.168.5.218:9999/api/orders';
    const payload = JSON.stringify({
        or_userid: "a65af5a5-d8e8-4a60-a881-e5e156891736",
        or_products: [
            {
                orp_productid: "fd0e5db8-decc-4be8-b76a-d552b205a7c5",
                orp_quantidad: 2,
                orp_totalprice: 5.0
            },
            {
                orp_productid: "ff0fa0e8-7f6d-42ca-8ca3-47be9e89b1c4",
                orp_quantidad: 2,
                orp_totalprice: 45.00
            }
        ],
        or_total_amount: 50.00,
        or_order_status: "pending",
        or_payment_method: "credit_card",
        or_shipping_address: "13dd1141-2001-4025-a058-8ea668735e82",
        or_created_at: "2024-07-30T12:34:56Z",
        or_updated_at:"2024-07-30T12:34:56Z"
    });

    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    http.post(url, payload, params);
}