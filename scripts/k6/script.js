//import GetOrders from "./scenarios/Get-Orders.js"
//import PostAddress from "./scenarios/Post-Address.js"
import PostUsers from "./scenarios/Post-Users.js"

import {group} from 'k6'

export default () => {
    group('Endpoint Get Orders', () => {
        //GetOrders()
        // PostAddress()
        PostUsers()
    })
}

//comando para iniciar: k6 run index.js --vus 50 --duration 60s --iterations 15000 --rps 250