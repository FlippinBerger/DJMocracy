import {defineStore} from 'pinia';

interface State {
    username: string
    userId: string
}

export const useUsersStore = defineStore({
    id: 'user-store',
    state: (): State => ({username: '', userId: ''}),
    actions: {
        logIn(username: string, userId: string) {
            this.username = username;
            this.userId = userId;
        },

        logOut() {
            this.username = '';
            this.userId = '';
        }
    },
    getters: {
        isLoggedIn: (state) => state.username !== '' && state.userId !== '',
    },
    persist: true,
})
