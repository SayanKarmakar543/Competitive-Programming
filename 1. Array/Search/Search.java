package Search;
import java.util.Scanner;

class Search{
    public static final Scanner scan = new Scanner(System.in);
    public static void main(String args[]){
        // int arr[] = new int[10];
        int arr[] = new int[]{10, 20, 30, 40, 50};
        System.out.print("Enter the serching element: ");
        int n = scan.nextInt();
        int pos = SearchingAnElement(arr, n);
        if (pos!=-1){
            System.out.println(pos);
        }
    }
    public static int SearchingAnElement(int arr[], int n){
        for (int i=0; i<arr.length; i++){
            if(arr[i]==n){
                return i;
            }
        }
        return -1;
    }
}