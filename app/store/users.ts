import {defineStore} from 'pinia';

export const useUsersStore = defineStore({
    id: 'user-store',
    state: () => ({username: '', token: ''}),
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
    persist: true,
})
