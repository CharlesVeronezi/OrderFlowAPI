import http from 'k6/http'
import { sleep } from 'k6'
import { Trend, Rate, Counter } from 'k6/metrics'
import { check, fail } from 'k6'

// export let GetOrdersDuration = new Trend('get_orders_duration')
// export let GetOrdersFailRate = new Rate('get_orders_fail_rate')
// export let GetOrdersSuccessRate = new Rate('get_orders_success_rate')
// export let GetOrdersReqs = new Rate('get_orders_reqs')

export default function () {
    http.put('http://192.168.5.218:9999/orders/ded06d27-419f-4727-8022-9fa3a9947c5d/conclude')

}