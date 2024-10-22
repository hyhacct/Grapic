<template>
    <div>


        <v-layout class="rounded rounded-md">
            <v-app-bar title="App">
            </v-app-bar>

            <v-navigation-drawer>
                <ViewForm />
            </v-navigation-drawer>

            <v-main>
                <ViewInput />
            </v-main>
        </v-layout>

    </div>
</template>

<script setup>
import ViewForm from "./v-form.vue"
import ViewInput from "./v-input.vue"

</script>