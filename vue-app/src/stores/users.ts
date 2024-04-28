import {defineStore} from 'pinia';

interface State {
    username: string
}

export const useUsersStore = defineStore({
    id: 'user-store',
    state: (): State => ({username: ''}),
    actions: {
        logIn(username: string) {
            this.username = username;
        },

        logOut() {
            this.username = '';
        }
    },
    getters: {
        isLoggedIn: (state) => state.username !== '',
    },
    persist: true,
})
