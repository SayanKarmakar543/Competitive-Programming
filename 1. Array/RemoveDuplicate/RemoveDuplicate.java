class RemoveDuplicate{
    public static void main(String[] args) {
        int arr[] = new int[]{10, 20, 20, 30, 30, 30, 30, 20};
        int brr[] = remDup(arr, arr.length);
        for (int i : brr) {
            System.out.print(i+" ");
        }
    }
    public static int[] remDup(int arr[], int n){
        int temp[] = new int[n];
        temp[0] = arr[0];
        int k = 1;
        for(int i=1;i<n;i++){
            boolean flag = true;
            for(int j=i-1;j>=0;j--){
                if(arr[i]==arr[j]){
                    // System.out.println("Sayan");
                    flag = false;
                    break;
                }
            }
            if(flag==true){
                temp[k++] = arr[i];
            }
        }
        int brr[] = new int[k];
        for(int i=0; i<k;i++){
            brr[i] = temp[i];
        }
        return brr;
    }
}