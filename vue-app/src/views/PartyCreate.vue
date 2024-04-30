<script setup>
import { ref } from 'vue';
import { useUsersStore } from '@/stores/users';

const partyName = ref("");
const partyTime = ref(new Date());

const store = useUsersStore();

const createParty = async () => {
    // type of partyTime.value is a string
    console.log(`partyName: ${partyName.value} partyTime: ${partyTime.value} partyTime type: ${typeof (partyTime.value)}`);

    const response = await fetch('http://localhost:1323/party', {
        body: JSON.stringify({
            name: partyName.value,
            time: partyTime.value,
            creator: store.username,
        }),
        headers: {
            "Content-Type": "application/json"
        },
        method: 'POST',
        credentials: 'include',
    })

    if (response.status === 401) {
        // TODO this is unauthorized error from server saying creds were
        // wrong
    }

    if (response.status !== 200) {
        throw new Error("Unable to login");
    }

    const data = await response.json();
}
</script>

<template>
    <div class='page'>
        <h1>Create Party</h1>
        <form class='form'>
            <input v-model="partyName" placeholder='Event name' />
            <input type="datetime-local" v-model="partyTime" />
            <button class='btn' @click="createParty">Create</button>
        </form>
    </div>
</template>

<style scoped>
.form {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.btn {
    align-self: center;
    width: 25%;
}
</style>
