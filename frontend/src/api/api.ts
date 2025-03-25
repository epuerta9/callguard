import { toast } from "@/hooks/use-toast";
import { useStore } from "@/context/store-context";

const host = "http://localhost:8000";

const headers = async () => {
  const result: { [key: string]: string } = {
    "Content-Type": "application/json",
    Accept: "application/json",
  };

  const userStr = localStorage.getItem("user");
  if (userStr) {
    try {
      const user = JSON.parse(userStr);
      if (user?.access_token) {
        result["Authorization"] = `Bearer ${user.access_token}`;
      }
    } catch {
      // If parsing fails, continue without token
    }
  }

  return result;
};
let pendingRequests = 0;

const useApi = () => {
  const { setLoader } = useStore();

  const incrementLoader = () => {
    if (pendingRequests === 0) {
      setLoader(true);
    }
    pendingRequests++;
  };

  const decrementLoader = () => {
    pendingRequests--;
    if (pendingRequests === 0) {
      setLoader(false);
    }
  };

  interface IConfig {
    loader?: boolean;
    toaster?: boolean;
  }

  const Api = {
    get: async (route: string, config?: IConfig) => {
      return Api.func(route, null, "GET", config);
    },
    put: async (route: string, params: any, config?: IConfig) => {
      return Api.func(route, params, "PUT", config);
    },
    post: async (route: string, params: any, config?: IConfig) => {
      return Api.func(route, params, "POST", config);
    },
    delete: async (route: string, params?: any, config?: IConfig) => {
      return Api.func(route, params, "DELETE", config);
    },
    patch: async (route: string, params: any, config?: IConfig) => {
      return Api.func(route, params, "PATCH", config);
    },

    async postImage(route: string, key: string, image: File, filetype: string) {
      const url = `${host}/${route}`;
      const formData = new FormData();
      const user = JSON.parse(localStorage.getItem("user") || "{}");

      formData.append(key, image);
      formData.append("description", "photo description");
      formData.append("file_name", image.name);
      formData.append("file_type", filetype || "products");
      formData.append("account_id", user.account_id);

      let options: RequestInit = {
        method: "POST",
        body: formData,
        headers: {},
      };

      if (user && user.access_token) {
        options.headers = {
          Authorization: `Bearer ${user.access_token}`,
        };
      }

      options = {
        method: "POST",
        body: formData,
        headers: options.headers,
      };

      try {
        const res = await fetch(url, options);
        if (!res.ok) throw new Error("Error posting image");
        const response = await res.json();
        return { data: response, error: undefined };
      } catch (error: any) {
        toast({
          title: "Error",
          description: "Failed to upload image",
          variant: "destructive"
        });
        console.error("Error posting image", error);
        return {
          data: undefined,
          error: error.response?.data || error.message,
        };
      }
    },

    async func(route: string, params: any, verb: string, config?: IConfig) {
      const { loader = true, toaster = true } = config || {};
      const url = `${host}${route}`;
      const options: any = {
        method: verb,
        headers: await headers(),
      };

      if (params) {
        options.body = JSON.stringify(params);
      }

      try {
        if (loader) incrementLoader();
        // console.log(verb, url, params || "");
        const res = await fetch(url, options);
        if (res.status === 204) return { data: "done", error: undefined };

        const response = await res.json();
        // console.log(verb, url, { response });

        if (res.status >= 200 && res.status < 300) {
          return { data: response, error: undefined };
        }

        if (res.status === 401) {
          localStorage.removeItem("access_token");
          localStorage.removeItem("user");
          window.location.href = "/sign-in";
          return { data: undefined, error: "Unauthorized" };
        }

        if (res.status === 400 && response.details) {
          const message = response.details.split(":")[0].replace(/"/g, "");
          toaster && toast({
            title: "Error",
            description: message,
            variant: "destructive"
          });
          return { data: undefined, error: message };
        }

        if (res.status === 500 && response.details) {
          const message =
            response.details
              .split(",")
              .find((part: string) => part.startsWith('"message":'))
              ?.replace(/("|:|message)/g, "") ?? response.details;
          if (message) {
            toaster && toast({
              title: "Error",
              description: message,
              variant: "destructive"
            });
            return { data: undefined, error: message };
          }
        }

        if (res.status === 403) return { data: null, error: "Forbidden" };

        if (res.status > 400) {
          const { error, status_code } = response;
          if (status_code === 401 && error === "Unauthorized") {
            localStorage.removeItem("user");
            return { data: undefined, error: "Unauthorized" };
          }
        }
      } catch (error: any) {
        console.error("Error:", error);
        toaster && toast({
          title: "Error",
          description: "An unexpected error occurred",
          variant: "destructive"
        });
        return {
          data: undefined,
          error: error.response?.data || error.message,
        };
      } finally {
        if (loader) decrementLoader();
      }
    },
  };

  return Api;
};

export default useApi;
