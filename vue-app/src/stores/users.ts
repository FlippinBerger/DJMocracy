import {defineStore} from 'pinia';

interface State {
    username: string
    token: string
}

export const useUsersStore = defineStore({
    id: 'user-store',
    state: (): State => ({username: '', token: ''}),
    actions: {
        logIn(username: string, token: string) {
            this.username = username;
            this.token = token;
        },

        logOut() {
            this.username = '';
            this.token = '';
        }
    },
    getters: {
        isLoggedIn: (state) => state.username !== '',
    },
    // persist: true,
})
