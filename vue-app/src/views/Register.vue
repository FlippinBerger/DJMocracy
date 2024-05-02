<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useUsersStore } from '@/stores/users';

const username = ref("");
const password = ref("");

const store = useUsersStore();

const router = useRouter();

const register = async () => {
    try {
        const response = await fetch('http://localhost:1323/signup', {
            body: JSON.stringify({
                username: username.value,
                password: password.value,
            }),
            headers: {
                "Content-Type": "application/json"
            },
            method: 'POST',
        })

        // TODO handle user already exists
        if (response.status !== 200) {
            throw new Error("Couldn't log in");
        }

        const data = await response.json();
        store.logIn(data.username, data.userId);

        router.push({ name: 'home' });
    } catch (err) {
        console.log('error logging in', err);
    }
}
</script>

<template>
    <div class='page'>
        <form class='form'>
            <input placeholder='Username' v-model="username"></input>
            <input placeholder='Password' type='password' v-model="password"></input>
            <button class='btn' type="button" @click='register'>Sign Up</button>
        </form>
    </div>
</template>

<style scoped>
.title {
    margin: 16px;
}

.form {
    display: flex;
    flex-direction: column;
    gap: 4px;
}

input {
    padding: 4px;
    min-width: 20%;
}

.btn {
    padding: 4px;
    align-self: center;
    width: 60%;
    margin-top: 8px;
}
</style>
