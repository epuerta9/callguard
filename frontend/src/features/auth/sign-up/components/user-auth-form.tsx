import { HTMLAttributes, useState } from "react";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate } from "@tanstack/react-router";
import {
  IconBrandFacebook,
  IconBrandGithub,
  IconCircleCheck,
} from "@tabler/icons-react";
import { cn } from "@/lib/utils";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { PasswordInput } from "@/components/password-input";
import { toast } from "@/hooks/use-toast";
import useApi from "@/api/api";
import { useAuth } from "@/hooks/use-auth";

type UserAuthFormProps = HTMLAttributes<HTMLDivElement>;

const formSchema = z.object({
  name: z
    .string()
    .min(1, { message: "Please enter your name" }),
  email: z
    .string()
    .min(1, { message: "Please enter your email" })
    .email({ message: "Invalid email address" }),
  password: z
    .string()
    .min(1, { message: "Please enter your password" })
    .min(8, { message: "Password must be at least 8 characters long" }),
});

export function UserAuthForm({ className, ...props }: UserAuthFormProps) {
  const [isLoading, setIsLoading] = useState(false);
  const navigate = useNavigate();
  const api = useApi();
  const { setUser } = useAuth();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    setIsLoading(true);

    try {
      const response = await api.post("/signup", values);
      
      if (!response) {
        throw new Error("No response from server");
      }

      if (response.error) {
        throw new Error(typeof response.error === 'string' ? response.error : 'Registration failed');
      }

      if (response.data) {
        setUser(response.data);
        toast({
          title: "Account created successfully!",
          description: "Welcome to CallGuard",
        });
        navigate({ to: "/" });
      }
    } catch (error) {
      toast({
        title: "Error",
        description: "Something went wrong. Please try again.",
        variant: "destructive",
      });
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <div className={cn("grid gap-6", className)} {...props}>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Name</FormLabel>
                <FormControl>
                  <Input
                    id="name"
                    placeholder="John Doe"
                    type="text"
                    disabled={isLoading}
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input
                    id="email"
                    placeholder="name@example.com"
                    type="email"
                    autoCapitalize="none"
                    autoComplete="email"
                    autoCorrect="off"
                    disabled={isLoading}
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <PasswordInput
                    id="password"
                    disabled={isLoading}
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button className="w-full" type="submit" disabled={isLoading}>
            {isLoading && (
              <IconCircleCheck className="mr-2 h-4 w-4 animate-spin" />
            )}
            Create Account
          </Button>
        </form>
      </Form>
      <div className="relative">
        <div className="absolute inset-0 flex items-center">
          <span className="w-full border-t" />
        </div>
        <div className="relative flex justify-center text-xs uppercase">
          <span className="bg-background px-2 text-muted-foreground">
            Or continue with
          </span>
        </div>
      </div>
      <div className="grid grid-cols-2 gap-6">
        <Button variant="outline" type="button" disabled={isLoading}>
          <IconBrandGithub className="mr-2 h-4 w-4" />
          Github
        </Button>
        <Button variant="outline" type="button" disabled={isLoading}>
          <IconBrandFacebook className="mr-2 h-4 w-4" />
          Facebook
        </Button>
      </div>
    </div>
  );
}
