import http from 'k6/http'
import { sleep } from 'k6'
import { Trend, Rate, Counter } from 'k6/metrics'
import { check, fail } from 'k6'

export let GetOrdersDuration = new Trend('get_orders_duration')
export let GetOrdersFailRate = new Rate('get_orders_fail_rate')
export let GetOrdersSuccessRate = new Rate('get_orders_success_rate')
export let GetOrdersReqs = new Rate('get_orders_reqs')

export default function () {
    let res = http.get('http://192.168.5.218:9999/orders/9b7b4ded-1efe-4c58-a9b1-0e88b2984948')
    //let res = http.get('http://localhost:9999/orders')

    GetOrdersDuration.add(res.timings.duration)
    GetOrdersReqs.add(1)
    GetOrdersFailRate.add(res.status == 0 || res.status > 300)
    GetOrdersSuccessRate.add(res.status < 300)

    let durationMsg = 'Max Duration ${500/1000}s'
    if (!check(res, {
        'max duration': (r) => r.timings.duration < 50000,
    })) {
        fail(durationMsg)
    }
}