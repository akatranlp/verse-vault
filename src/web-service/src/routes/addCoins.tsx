import { Button } from "@/components/ui/button";
import { useUserData } from "@/provider/user-provider.tsx";
import { addCoins } from "@/repository/user.ts";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import toast from "react-hot-toast";

export const AddCoins = () => {
  const user = useUserData();
  const queryClient = useQueryClient();

  const { mutate, variables, isPending } = useMutation({
    mutationFn: (updateUser: UpdateUser) => addCoins(updateUser),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["me"] });
      toast.success(`New Balance ${variables?.balance}VV-Coins`);
    },
    onError: () => {
      toast.error("An error occurred. Please try again.");
    },
  });

  const handleAddCoinsClick = async (amount: number) => {
    const updatedUser = {
      balance: user.balance + amount,
    };
    mutate(updatedUser);
  };

  return (
    <>
      <div className="text-center text-4xl pt-2.5 mx-auto">Add VV-Coins</div>
      <div className="flex justify-center pt-5">
        <div className="grid grid-cols-6 justify-center items-center gap-4 w-[50%]">
          <Button className="col-span-2 w-full" variant="secondary" disabled={isPending} onClick={() => handleAddCoinsClick(100)}>
            100 VV-Coins
          </Button>
          <Button className="col-span-2 w-full" variant="secondary" disabled={isPending} onClick={() => handleAddCoinsClick(250)}>
            250 VV-Coins
          </Button>
          <Button className="col-span-2 w-full" variant="secondary" disabled={isPending} onClick={() => handleAddCoinsClick(500)}>
            500 VV-Coins
          </Button>
          <Button className="col-start-2 col-span-2 w-full" variant="secondary" disabled={isPending} onClick={() => handleAddCoinsClick(750)}>
            750 VV-Coins
          </Button>
          <Button className="col-span-2 w-full" variant="secondary" disabled={isPending} onClick={() => handleAddCoinsClick(1000)}>
            1000 VV-Coins
          </Button>
        </div>
      </div>
    </>
  );
};