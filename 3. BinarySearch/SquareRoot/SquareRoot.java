

class SquareRoot{
    public static void main(String[] args) {
        System.out.println(SqRoot(-8));
    }
    public static int SqRoot(int n){
        int low = 0;
        int high = n;
        int ans = -1;
        if(n<0){
            return ans;
        }
        while(low<=high){
            int mid = (low+high)/2;
            if(mid*mid==n){
                return mid;
            }
            else if(mid*mid<n){
                ans = mid;
                low=mid+1;
            }
            else{
              high = mid-1;  
            }
        }
        return ans;
    }
}