// import GetOrders from "./scenarios/Get-Orders.js"
//import PostAddress from "./scenarios/Post-Address.js"
//import PostUsers from "./scenarios/Post-Users.js"
import PostOrders from "./scenarios/Post-Orders.js"
//import PatchOrders from "./scenarios/Patch-Orders.js"

import {group} from 'k6'

export default () => {
    group('Endpoint Get Orders', () => {
        //GetOrders()
        // PostAddress()
        //PostUsers()
        PostOrders()
        //PatchOrders()

    })
}

//comando para iniciar: k6 run scripts/k6/index.js --vus 50 --duration 60s --iterations 15000 --rps 250