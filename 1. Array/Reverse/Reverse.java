import java.util.Scanner;

class Reverse{
    public static final Scanner scan = new Scanner(System.in);
    public static void main(String[] args) {
        int arr[] = new int[]{10, 20, 30, 40, 50};
        ReverseArray(arr);
        for (int i : arr) {
            System.out.print(i+" ");
        }
    }
    public static void ReverseArray(int[] arr){
        int low = 0;
        int high = arr.length-1;
        while(low<high){
            int temp = arr[low];
            arr[low] = arr[high];
            arr[high] = temp;
            low++;
            high--;
        }
        
    }
}