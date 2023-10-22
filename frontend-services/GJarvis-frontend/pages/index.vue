<template>
  <v-container class="ma-0">
    <v-row>
      <v-col class="pl-0">
          <v-layout>
            <v-navigation-drawer>
              <v-list>
                <v-list-item title="Jarvis" subtitle="Welcome to this page">
                </v-list-item>
              </v-list>

              <v-divider></v-divider>

              <v-list :lines="false" density="compact" nav>
                <v-list-item
                  v-for="(item, i) in items"
                  :key="i"
                  :value="item"
                  color="blue"
                >
                  <template v-slot:prepend>
                    <v-icon :icon="item.icon"></v-icon>
                  </template>

                  <v-list-item-title v-text="item.text"></v-list-item-title>
                </v-list-item>
              </v-list>
            </v-navigation-drawer>

            <v-main style="height: 354px"></v-main>
          </v-layout>
      </v-col>
      <v-col :cols="9">
        <div>
          <h1>aaa{{ count }}</h1>
          <h1>12312321312</h1>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup>
import { ref } from "vue";
const items = ref([
  { text: "My Files", icon: "mdi-folder" },
  { text: "Shared with me", icon: "mdi-account-multiple" },
  { text: "Starred", icon: "mdi-star" },
  { text: "Recent", icon: "mdi-history" },
  { text: "Offline", icon: "mdi-check-circle" },
  { text: "Uploads", icon: "mdi-upload" },
  { text: "Backups", icon: "mdi-cloud-upload" },
]);

const { data: count } = await useFetch(
  "/api/v1/voice-assistance-service/hello"
);
</script>

<!-- <script lang="ts" setup>
import { ref, onMounted } from "vue";
import axios from "axios";

const message = ref('');

onMounted(async () => {
    // try {
    //     const response = 
    await axios.get("http://localhost:8080/api/v1/voice-assistance-service/hello")
        .then((response) => message.value = response.data)
        .catch((error) => console.log(error))
    //     message.value = response.data;
    // } catch (error) {
    //     console.log(error);
    // }
}); -->
<!-- </script> -->
