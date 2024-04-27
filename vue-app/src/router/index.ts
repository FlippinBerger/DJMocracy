import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Party from '../views/Party.vue'
import PartyCreate from '../views/PartyCreate.vue'
import PartySearch from '../views/PartySearch.vue'
import Profile from '../views/Profile.vue'
import Register from '../views/Register.vue'
import Auth from '../views/Auth.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/auth',
            children: [
                { path: '', component: Auth, name: 'Auth'},
                { path: '/login', component: Login, name: 'Login' },
                { path: '/register', component: Register, name: 'Register' },
            ]
        },
        {
            path: '/parties/:partyId',
            name: 'party',
            component: Party,
        },
        {
            path: '/party-create',
            name: 'party-create',
            component: PartyCreate,
        },
        {
            path: '/party-search',
            name: 'party-search',
            component: PartySearch,
        },
        {
            path: '/profile',
            name: 'profile',
            component: Profile,
        },
        {
            path: '/',
            name: 'home',
            component: Home
        },
    ]
})

export default router
