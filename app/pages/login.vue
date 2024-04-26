<script setup>
    import {useUsersStore} from '~/store/users';

    const username = ref("");
    const password = ref("");

    const store = useUsersStore();

    const logIn = async () => {
        store.logIn(username.value, password.value);
        // navigateTo('/');
    }
</script>

<template>
    <div class='page'>
        <form class='form'>
            <input placeholder='Username' v-model="username"></input>
            <input placeholder='Password' type='password' v-model="password"></input>
            <!-- TODO maybe make this a button again if we can figure out the refresh problem -->
            <!-- <button class='btn' type=submit @click='logIn'>Log in</button> -->
            <NuxtLink to='/' @click="logIn">Log in</NuxtLink>
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
