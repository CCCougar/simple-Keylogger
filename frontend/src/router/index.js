// 配置路由相关信息
import VueRouter from 'vue-router'
import Vue from "vue"
// import showBeacons from "@/components/ShowBeacons";
import beaconDetails from "@/components/BeaconDetails";

Vue.use(VueRouter)

const routes = [
    // {
    //     path:'/',
    //     component: showBeacons
    // },
    {
        path:'/:uuid',
        component: beaconDetails
    }
]

const router = new VueRouter({
    routes,
    mode:'history'
})

// router.beforeEach((to, from, next) => {
//
//     next()
// })

// router.afterEach((to, from) => {
//     console.log(to)
//     console.log(from)
// })

export default router