
import java.util.Scanner;

class CheckSortedArray {
    public static final Scanner scan = new Scanner(System.in);
    public static void main(String[] args) {
        int arr[] = new int[]{10, 20, 30, 40, 50};
        boolean status = CheckSortedArrayMethod(arr);
        System.out.print(status);
    }
    public static boolean CheckSortedArrayMethod(int arr[]){
        boolean status = true;
        for(int i=1; i<arr.length; i++){
            if(arr[i-1]>arr[i]){
                status = false;
                break;
            }
        }
        return status;
    }
}
