package Largest;
import java.util.Scanner;

public class Largest {
    public static final Scanner scan = new Scanner(System.in);
    public static void main(String[] args) {
        int arr[] = new int[]{10, 20, 30, 40, 50};
        int ans = FindLargestElement(arr);
        System.out.println(ans);
    }
    public static int FindLargestElement(int arr[]){
        int max = arr[0];
        for (int i=0; i<arr.length; i++){
            if(arr[i] > max) {
                max = arr[i];
            }
        }
        return max;
    }
}
