<script setup lang="ts">
import { reactive } from "vue";
import { useRouter } from "vue-router";
import { useStore } from "@/stores";

const data = reactive({
  email: "",
  password: "",
});

const router = useRouter();

const submit = async () => {
  try {
    const response = await fetch("http://localhost:5000/signin", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify(data),
    });

    if (response.status !== 200) {
      useStore().isAuthenticated = false
      router.push("/signin");
      return
    } else {
      useStore().isAuthenticated = true
      router.push("/");
      return
    }
  } catch (e) {
    useStore().isAuthenticated = false
    return
  }
};
</script>

<template>
  <section class="bg-gray-900">
    <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
      <div class="w-full rounded-lg shadow border md:mt-0 sm:max-w-md xl:p-0 bg-gray-800 border-gray-700">
        <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 class="text-xl font-bold leading-tight tracking-tight md:text-2xl text-white text-center">
            Sign in to your account
          </h1>
          <form @submit.prevent="submit" class="space-y-4 md:space-y-6">
            <div>
              <label for="email" class="block mb-2 text-sm font-medium text-white">Your Email</label>
              <input v-model="data.email" type="email" name="email" id="email" class="border sm:text-sm focus:outline-none focus:ring-2 rounded-lg block w-full p-2.5 bg-gray-700 border-gray-600 placeholder-gray-400 text-white focus:ring-blue-800 focus:border-blue-800" placeholder="name@company.com" required />
            </div>
            <div>
              <label for="password" class="block mb-2 text-sm font-medium text-white">Password</label>
              <input v-model="data.password" type="password" name="password" id="password" placeholder="••••••••" class="border sm:text-sm focus:outline-none focus:ring-2 rounded-lg block w-full p-2.5 bg-gray-700 border-gray-600 placeholder-gray-400 text-white focus:ring-blue-800 focus:border-blue-800" required />
            </div>
            <button type="submit" class="w-full text-white focus:ring-2 focus:outline-none font-medium rounded-lg text-sm px-5 py-2.5 text-center bg-blue-600 hover:bg-blue-700 focus:ring-blue-800">
              Sign in
            </button>
            <p class="text-sm font-light text-gray-400">
              Don’t have an account yet?
              <RouterLink to="/signup" class="font-medium hover:underline text-blue-500">Sign up</RouterLink>
            </p>
          </form>
        </div>
      </div>
    </div>
  </section>
</template>
